package storage

import (
	"context"
	"fmt"
	"strconv"

	sq "github.com/Masterminds/squirrel"
	"github.com/njslxve/tender-service/internal/entity"
)

func (s *Storage) GetBidReviews(authorUsername string, limit string, offset string) ([]entity.BidFeedback, error) {
	const op = "storage.GetBidReviews"

	var lim, off int

	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		lim = l
	}

	if offset != "" {
		o, err := strconv.Atoi(offset)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		off = o
	}

	querry := qb.Select("id", "description", "created_at").
		From("bid_feedback bf").
		Join("bids b ON bf.bid_id = b.id").
		Join("employee e ON b.author_id = e.id").
		Where(sq.Eq{"e.username": authorUsername})

	if limit != "" {
		querry = querry.Limit(uint64(lim))
	}

	if offset != "" {
		querry = querry.Offset(uint64(off))
	}

	sql, args, err := querry.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var feedbacks []entity.BidFeedback

	rows, err := s.db.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		var feedback entity.BidFeedback
		err = rows.Scan(&feedback.ID, &feedback.Description, &feedback.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		feedbacks = append(feedbacks, feedback)
	}

	return feedbacks, nil
}
